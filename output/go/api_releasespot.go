/*
 * URx robot controller API docs
 *
 * API for the URx robot controller
 *
 * API version: 1.0.0
 * Contact: danai.vachtsevanou@unisg.ch
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"net/http"
)

func ReleasespotPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
