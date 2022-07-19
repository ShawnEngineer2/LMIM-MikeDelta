package handlers

func GenericSvcErrHandler(newErr error) string {
	if newErr != nil {
		return newErr.Error()
	} else {
		return ""
	}
}
