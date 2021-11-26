package yolo

import (
	"github.com/justtaldevelops/go-hcaptcha/solver"
	"github.com/sirupsen/logrus"
	"github.com/wimspaargaren/yolov3"
	"gocv.io/x/gocv"
	"strings"
)

// yolo is the YOLO v3 network.
var yolo yolov3.Net

// init initializes the YOLO v3 network.
func init() {
	yolo, _ = yolov3.NewNet("yolo/yolov3.weights", "yolo/yolov3.cfg", "yolo/coco.names")
}

// YOLOSolver uses the "You Only Look Once" (YOLO) algorithm to solve hCaptcha tasks.
type YOLOSolver struct {
	// Log is the logger for the solver.
	Log *logrus.Logger
}

// Solve ...
func (s *YOLOSolver) Solve(category, object string, tasks []solver.Task) []solver.Task {
	// Make sure the YOLO network is initialized.
	if yolo == nil {
		panic("yolov3 data is not in expected folders")
	}

	// Make sure we can solve the challenge.
	if category != "image_label_binary" {
		s.Log.Debugf("cannot solve challenge with category %s", category)
		return []solver.Task{}
	}

	// Answer the challenge.
	var answers []solver.Task
	for _, task := range tasks {
		// Decode and detect the object.
		frame, err := gocv.IMDecode(task.Image, gocv.IMReadColor)
		if err != nil {
			continue
		}

		detections, err := yolo.GetDetections(frame)
		if err != nil {
			continue
		}

		for _, detection := range detections {
			fixedClassName := strings.TrimSpace(detection.ClassName)
			if fixedClassName == object && detection.Confidence > 0.6 {
				s.Log.Debugf("Detected %v in provided image", object)

				answers = append(answers, task)
				break
			}
		}
	}

	return answers
}
