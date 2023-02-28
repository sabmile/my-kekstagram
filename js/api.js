import { URL_DATA } from "./const.js"

const getData = (onSuccess) => {
    fetch(URL_DATA)
        .then((response) => {
            if (response.ok) {
                return response.json();
            }
        })
        .then((pictures) => {
            onSuccess(pictures);
        })
}

export { getData }