import cv2
import sys
import os

# Get user supplied values
directory = sys.argv[1]
cascPath = "haarcascade_frontalface_default.xml"

# Create the haar cascade
faceCascade = cv2.CascadeClassifier(cascPath)

def isFaceFound(image):

    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

    return faceCascade.detectMultiScale(
        gray,
        scaleFactor=1.1,
        minNeighbors=5,
        minSize=(30, 30))

def notify(filename, faces, image):
    if len(faces) > 0:

        print(filename + ": faces found: {0}".format(len(faces)))

        # draw a rectangle around the faces
        for (x, y, w, h) in faces:
            cv2.rectangle(image, (x, y), (x+w, y+h), (0, 255, 0), 2)

        cv2.imshow("faces found", image)
        cv2.waitKey(13)
        raw_input("next?")

    else:
        print(filename + ": did not find faces")

def checkDir(imageDir):
    for filename in os.listdir(imageDir):
        if filename.endswith(".jpg"):
            path = imageDir + "/" + filename
            image = cv2.imread(path)
            faces = isFaceFound(image)
            notify(path, faces, image)
            continue
        else:
            continue

checkDir(directory)
