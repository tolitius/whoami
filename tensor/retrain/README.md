label images with a (new) retrained model:

```bash
python label_image.py --graph=/Users/tolitius/1/playground/tensor/retrain/models/humans/output_graph.pb \
                      --labels=/Users/tolitius/1/playground/tensor/retrain/models/humans/output_labels.txt \
                      --image=path-to-image.jpg
```
```bash
human (score = 0.63270)
face (score = 0.31476)
hat (score = 0.05017)
glasses (score = 0.00237)
```
