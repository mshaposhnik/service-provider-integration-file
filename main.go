// Copyright (c) 2021 Red Hat, Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	start()
}
func OkHandler2(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func OkHandler(w http.ResponseWriter, r *http.Request) {
	//repo := r.FormValue("repo")
	//path := r.FormValue("path")
	//getter.Get()
	//contents, err := getFileContents(context.TODO(), repo, path, "", func(url string) {
	//	http.Redirect(w, r, url, http.StatusFound)
	//})
	//
	//if err != nil {
	//	return
	//}
	//_, _ = io.Copy(w, contents)
}

func start() {
	router := mux.NewRouter()

	router.HandleFunc("/", OkHandler2).Methods("GET")

	err := http.ListenAndServe(fmt.Sprintf(":%d", 8000), router)
	if err != nil {
		zap.L().Error("failed to start the HTTP server", zap.Error(err))
	}
}
