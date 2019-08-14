package rest

import (
	"net/http"

	"github.com/LINBIT/linstor-remote-storage/iscsi"
)

func ISCSICreate(w http.ResponseWriter, r *http.Request) {
	var iscsiCfg iscsi.ISCSI
	if err := unmarshalBody(w, r, &iscsiCfg); err != nil {
		return
	}

	if err := iscsiCfg.CreateResource(); err != nil {
		Errorf(http.StatusInternalServerError, w, "Could not create resource: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(xxx)
}