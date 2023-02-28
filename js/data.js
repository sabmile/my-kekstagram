import {
    getRandomPositiveInteger,
    createIdGenerator,
} from "./util.js"

import {
    ID,
    COUNT_PICTURES,
    MIN_COUNT_PICTURES,
    COUNT_COMMENTS,
    MIN_COUNT_COMMENTS,
    COUNT_AVATARS,
    MIN_AVATAR,
    MAX_AVATAR,
    MIN_COUNT_LIKES,
    MAX_COUNT_LIKES,
} from "./const.js"

const NAMES = [
    "Elisa Rush",
    "Lulu Estes",
    "Brogan Atkins",
    "Alessandro Alvarez",
    "Jerome Briggs",
    "Maliha Petersen",
    "Cordelia Barrett",
    "Lexie Castillo",
    "Harri Stevens",
    "Veronica Ross",
]

const COMMENTS = [
    "When you take a photo, it would be good to remove your finger from the frame. In the end, it's just unprofessional. One thing is unclear: how so?!",
    "There is simply no framing. The filter was selected unsuccessfully: I would use a series set to 80%",
    "I lost my family, children and cat because of this photo. They said they didn't share my love for art and went to a neighbor. Everything is fine!",
    "I'm stuck on this photo and I can't tear myself away. I don't know what to do at all. Is this a composition?! What is this composition?!",
    "I slipped on a banana peel and dropped the camera on the cat and I got a better picture.",
    "The faces of the people in the photo are distorted, as if they are being beaten. How could you catch such an unfortunate moment?! The horizon is littered.",
    "My grandmother accidentally sneezed with a camera in her hands and she got a better photo. Shob I lived like this!",
    "I can't imagine how you can photograph the sea and sunset better. This is just the apogee. After that, we can burn all the cameras, because the peak has been reached anyway. The focus is blurred. Or is it just someone splattered the lens?",
]

const DESCRIPTIONS = [
    "If you clearly formulate a wish for the universe, then everything will definitely come true. Believe in yourself. The main thing is to want and dream.",
    "Appreciate every moment. Appreciate those who are close to you and drive away all doubts. Don't offend everyone with words",
    "How cool is the food here",
    "Hanging out with friends at the sea",
]

const getRandomArrayElement = (elements) => {
    return elements[getRandomPositiveInteger(0, elements.length - 1)];
};


const generateCommentId = createIdGenerator();
const generatePostId = createIdGenerator();

const createComment = () => {
    let id = generateCommentId()
    return {
        id,
        avatar: `img/avatar-${getRandomPositiveInteger(MIN_AVATAR, MAX_AVATAR)}`,
        message: getRandomArrayElement(COMMENTS),
        name: getRandomArrayElement(NAMES),
    }
}

const createPicture = () => {
    let id = generatePostId()
    return {
        id,
        url: `photos/${id}.jpg`,
        likes: getRandomPositiveInteger(MIN_COUNT_LIKES, MAX_COUNT_LIKES),
        comments: Array.from({ length: getRandomPositiveInteger(MIN_COUNT_COMMENTS, COUNT_COMMENTS) }, createComment),
        description: getRandomArrayElement(DESCRIPTIONS),
    }
}

const data = Array.from({ length: COUNT_PICTURES }, createPicture)

export { data }
