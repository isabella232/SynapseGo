package wrapper

/********** GLOBAL VARIABLES **********/
const instiURL = _url + "/institutions"

/********** METHODS **********/

// GetInstitutions returns all of the nodes associated with a user
func (c *ClientCredentials) GetInstitutions(userID string) map[string]interface{} {
	res, body, errs := request.
		Get(instiURL).
		EndBytes()

	if res != nil && errs != nil {
		errorLog(errs)
	}

	return format(body)
}
