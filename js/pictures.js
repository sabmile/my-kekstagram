const picturesBloc = document.querySelector('.pictures')
const pictureTemplate = document.querySelector('#picture').content.querySelector('.picture')

const renderPictures = (pictures) => {
    const fragment = document.createDocumentFragment();

    pictures.forEach((picture) => {
        const pic = pictureTemplate.cloneNode(true);

        pic.querySelector('.picture__img').src = picture.url;
        fragment.appendChild(pic)
    });

    picturesBloc.appendChild(fragment)
};

export { renderPictures }

