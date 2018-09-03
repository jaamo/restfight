/**
 * Helper functions to do requests.
 */
function doFetch(url, json) {
	let params = '?';
	for (var i in json) {
		params += i + '=' + json[i] + "&";
	}
	return fetch(url + params).then(r => r.json());
}

/**
 * Get game status.
 */
export function getStatus() {
	return dispatch => {
		return doFetch('http://localhost:8000/status')
			.then(json => console.log(json));
	}
}

/**
 * Show controls.
 */
export const showControls = () => ({
	type: 'SHOW_CONTROLS',
});

