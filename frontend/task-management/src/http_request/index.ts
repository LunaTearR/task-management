import createHttpRequest from "./createHttpRequest";

export const mainApiUrl: string = 'http://localhost:3000/api'

const httpRequestInstance = createHttpRequest(mainApiUrl);

export default httpRequestInstance;