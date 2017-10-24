package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/julienschmidt/httprouter"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

type ClassifyResult struct {
	Filename string        `json:"filename"`
	Labels   []LabelResult `json:"labels"`
}

type LabelResult struct {
	Label       string  `json:"label"`
	Probability float32 `json:"probability"`
}

type ByProbability []LabelResult

func (a ByProbability) Len() int           { return len(a) }
func (a ByProbability) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByProbability) Less(i, j int) bool { return a[i].Probability > a[j].Probability }

func findBestLabels(labels []string, probabilities []float32) []LabelResult {
	// Make a list of label/probability pairs
	var resultLabels []LabelResult
	for i, p := range probabilities {
		if i >= len(labels) {
			break
		}
		resultLabels = append(resultLabels, LabelResult{Label: labels[i], Probability: p})
	}
	// Sort by probability
	sort.Sort(ByProbability(resultLabels))
	// Return top 5 labels
	return resultLabels[:5]
}

func loadModel(graphPath, labelsPath string) (*tf.Graph, []string, error) {

	// loading inception model

	model, err := ioutil.ReadFile(graphPath)
	if err != nil {
		return nil, nil, err
	}
	graph := tf.NewGraph()
	if err := graph.Import(model, ""); err != nil {
		return nil, nil, err
	}

	// loading labels

	var labels []string

	labelsFile, err := os.Open(labelsPath)
	if err != nil {
		return nil, nil, err
	}
	defer labelsFile.Close()
	scanner := bufio.NewScanner(labelsFile)

	// labels are separated by newlines

	for scanner.Scan() {
		labels = append(labels, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return graph, labels, nil
}

func responseError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func makeTransformImageGraph(imageFormat string) (graph *tf.Graph, input, output tf.Output, err error) {
	const (
		H, W  = 224, 224
		Mean  = float32(117)
		Scale = float32(1)
	)
	s := op.NewScope()
	input = op.Placeholder(s, tf.String)

	// decoding png and jpeg

	var decode tf.Output
	if imageFormat == "png" {
		decode = op.DecodePng(s, input, op.DecodePngChannels(3))
	} else {
		decode = op.DecodeJpeg(s, input, op.DecodeJpegChannels(3))
	}

	// Div and Sub perform (value-Mean)/Scale for each pixel

	output = op.Div(s,
		op.Sub(s,
			// resize to 224x224 with bilinear interpolation
			op.ResizeBilinear(s,
				// create a batch containing a single image
				op.ExpandDims(s,
					// use decoded pixel values
					op.Cast(s, decode, tf.Float),
					op.Const(s.SubScope("make_batch"), int32(0))),
				op.Const(s.SubScope("size"), []int32{H, W})),
			op.Const(s.SubScope("mean"), Mean)),
		op.Const(s.SubScope("scale"), Scale))

	graph, err = s.Finalize()

	return graph, input, output, err
}

func makeTensorFromImage(imageBuffer *bytes.Buffer, imageFormat string) (*tf.Tensor, error) {
	tensor, err := tf.NewTensor(imageBuffer.String())
	if err != nil {
		return nil, err
	}
	graph, input, output, err := makeTransformImageGraph(imageFormat)
	if err != nil {
		return nil, err
	}
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	normalized, err := session.Run(
		map[tf.Output]*tf.Tensor{input: tensor},
		[]tf.Output{output},
		nil)
	if err != nil {
		return nil, err
	}
	return normalized[0], nil
}

func recognize(graph *tf.Graph, labels []string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Read image
		imageFile, header, err := r.FormFile("image")
		// Will contain filename and extension
		imageName := strings.Split(header.Filename, ".")
		if err != nil {
			responseError(w, "Could not read image", http.StatusBadRequest)
			return
		}
		defer imageFile.Close()
		var imageBuffer bytes.Buffer

		// Copy image data to a buffer
		io.Copy(&imageBuffer, imageFile)

		tensor, err := makeTensorFromImage(&imageBuffer, imageName[:1][0])
		if err != nil {
			responseError(w, "Invalid image", http.StatusBadRequest)
			return
		}

		session, err := tf.NewSession(graph, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()
		output, err := session.Run(
			map[tf.Output]*tf.Tensor{
				graph.Operation("input").Output(0): tensor,
			},
			[]tf.Output{
				graph.Operation("output").Output(0),
			},
			nil)
		if err != nil {
			responseError(w, "Could not run inference", http.StatusInternalServerError)
			return
		}
		responseJSON(w, ClassifyResult{
			Filename: header.Filename,
			Labels:   findBestLabels(labels, output[0].Value().([][]float32)[0]),
		})
	}
}

func main() {

	port := ":4242"
	modelPath := "/model/tensorflow_inception_graph.pb"
	labelsPath := "/model/imagenet_comp_graph_label_strings.txt"

	log.Println("raspberry pi inspector is warmig up..")

	graph, labels, err := loadModel(modelPath, labelsPath)

	if err != nil {
		log.Fatal("could not create a graph")
		return
	}

	log.Println("raspberry pi inspector: loaded models")

	r := httprouter.New()
	r.POST("/whoami", recognize(graph, labels))

	log.Printf("raspberry pi inspector: starting up on :%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
