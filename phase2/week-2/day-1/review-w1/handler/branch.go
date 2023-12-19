package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"game-store/entity"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateBranch handles POST Branch endpoint
func (h *Handler) CreateBranch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newBranch entity.Branch
	if err := json.NewDecoder(r.Body).Decode(&newBranch); err != nil {
		httpError(w, http.StatusBadRequest, err)
		return
	}

	ok := validateInput(newBranch)
	if !ok {
		httpError(w, http.StatusBadRequest, fmt.Errorf("payload body should consists of name, email, and phone"))
		return
	}

	Branch, err := h.branch.Create(newBranch)
	if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	httpSuccess(w, http.StatusCreated, Branch)
}

// GetBranch handles GET Branch by id endpoint
func (h *Handler) GetBranch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	BranchID := ps.ByName("id")
	Branch, err := h.branch.Get(BranchID)
	if err != nil && err == sql.ErrNoRows {
		httpError(w, http.StatusNotFound, fmt.Errorf("branch Id %v not found", BranchID))
		return
	} else if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	httpSuccess(w, http.StatusOK, Branch)
}

// ListBranch handles GET list Branch endpoint
func (h *Handler) ListBranch(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Branchs, err := h.branch.List()
	if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	httpSuccess(w, http.StatusOK, Branchs)
}

// UpdateBranch handles PUT Branch endpoint
func (h *Handler) UpdateBranch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	BranchID := ps.ByName("id")
	var newBranch entity.Branch
	if err := json.NewDecoder(r.Body).Decode(&newBranch); err != nil {
		httpError(w, http.StatusBadRequest, err)
		return
	}

	ok := validateInput(newBranch)
	if !ok {
		httpError(w, http.StatusBadRequest, fmt.Errorf("payload body should consists of name, email, and phone"))
		return
	}

	Branch, err := h.branch.Update(BranchID, newBranch)
	if err != nil && err == sql.ErrNoRows {
		httpError(w, http.StatusNotFound, fmt.Errorf("branch id %v doesn't exist", BranchID))
		return
	} else if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	httpSuccess(w, http.StatusAccepted, Branch)
}

// DeleteBranch handles DELETE Branch endpoint
func (h *Handler) DeleteBranch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	BranchID := ps.ByName("id")
	err := h.branch.Delete(BranchID)
	if err != nil && err == sql.ErrNoRows {
		httpError(w, http.StatusNotFound, fmt.Errorf("branch id %v doesn't exist", BranchID))
		return
	} else if err != nil {
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	httpSuccess(w, http.StatusOK, nil)
}

// validateInput prevents empty payload to be submitted
func validateInput(Branch entity.Branch) bool {
	if Branch.Name == "" {
		return false
	}

	if Branch.Location == "" {
		return false
	}

	return true
}
