package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/muchrm/golang-img-merge/imgmerge"
)

var (
	first  = []string{}
	second = []string{}
)

func showIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html>
    <html>
    
    <head>
        <meta http-equiv="content-type" content="text/html; charset=UTF-8">
        <script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/angularjs/1.4.4/angular.js"></script>
        <script type="text/javascript" src="https://angular-file-upload.appspot.com/js/ng-file-upload-shim.js"></script>
        <script type="text/javascript" src="https://angular-file-upload.appspot.com/js/ng-file-upload.js"></script>
        <script type='text/javascript'>
            var app = angular.module('fileUpload', ['ngFileUpload']);
            app.controller('MyCtrl', ['$scope', 'Upload', '$timeout', '$http', function ($scope, Upload, $timeout, $http) {
                $scope.uploadFiles = function (files, files2) {
                    $scope.files = files;
                    if (files && files.length) {
                        Upload.upload({
                            url: '/first',
                            data: {
                                files: files
                            }
                        }).then(function (response) {
                            Upload.upload({
                                url: '/second',
                                data: {
                                    files: files2
                                }
                            }).then(function (response) {
                                $http.get("/run")
                                    .then(function (response) {
                                        alert(response.data);
                                    });
                            });
                        })
                    }
                };
            }]);
        </script>
    </head>
    
    <body>
    
        <body ng-app="fileUpload" ng-controller="MyCtrl">
            <form name="myForm">
                <fieldset>
                    <legend>Upload on form submit</legend>
    
                    <br>เลือกลาย:
                    <input type="file" ngf-select ng-model="$files" multiple name="file" accept="image/*" ngf-max-size="2MB" required ngf-model-invalid="errorFile"
                    />
                    <br>เลือกเสื้อ:
                    <input type="file" ngf-select ng-model="$files2" multiple name="file" accept="image/*" ngf-max-size="2MB" required ngf-model-invalid="errorFile"
                    />
                    <button ng-disabled="!myForm.$valid" ng-click="uploadFiles($files,$files2)">Submit</button>
                </fieldset>
                <br>
            </form>
        </body>
    </body>
    
    </html>`)
}
func uploadFirst(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//copy each part to destination.
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		//if part.FileName() is empty, skip this iteration.
		if part.FileName() == "" {
			continue
		}
		path := "first/"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0700)
		}
		dst, err := os.Create(path + part.FileName())
		defer dst.Close()

		if err != nil {
			http.Error(w, "1"+err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, "2"+err.Error(), http.StatusInternalServerError)
			return
		}
		first = append(first, part.FileName())
	}

}
func uploadSecond(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//copy each part to destination.
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		//if part.FileName() is empty, skip this iteration.
		if part.FileName() == "" {
			continue
		}
		path := "second/"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0700)
		}

		dst, err := os.Create(path + part.FileName())
		defer dst.Close()

		if err != nil {
			http.Error(w, "1"+err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, "2"+err.Error(), http.StatusInternalServerError)
			return
		}
		second = append(second, part.FileName())
	}
}
func runImgMerge(w http.ResponseWriter, r *http.Request) {
	path := "out/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700)
	}
	for _, img := range first {
		for _, img2 := range second {
			imgmerge.MergeImage("first/"+img, "second/"+img2, fmt.Sprintf("%s%s-%s", path, img, img2))
		}
	}
	fmt.Fprintf(w, "เสร็จแล้วจ้า") // send data to client side
}

func main() {
	http.HandleFunc("/", showIndex)
	http.HandleFunc("/first", uploadFirst)
	http.HandleFunc("/second", uploadSecond)
	http.HandleFunc("/run", runImgMerge)     // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
