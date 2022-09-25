import axios from 'axios'

export default function send(url: string, params: Object) {
  return axios.post(url, params)
}
