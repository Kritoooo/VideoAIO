import axios from 'axios'

const baseURL = '/api/v1'

const api = axios.create({
  baseURL,
  timeout: 60000
})

export const uploadVideo = (file: File) => {
  const formData = new FormData()
  formData.append('video', file)
  return api.post('/video/upload', formData)
}

export const processVideo = (videoId: string) => {
  return api.post('/video/process', { videoId })
}
