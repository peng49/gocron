import httpClient from '../utils/httpClient'

export default {
    list(query, callback) {
        httpClient.get('/permission/roles', query, callback)
    },
    store(data, callback) {
        httpClient.post('/permission/role/store', data, callback)
    },
}