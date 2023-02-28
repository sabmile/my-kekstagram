import { renderPictures } from "./pictures.js"
import { getData } from "./api.js";

getData((data) => {
    renderPictures(data)
});

