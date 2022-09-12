import axios, { AxiosResponse } from 'axios'
// import { useStore } from '@/store/index'
// const store = useStore()

// interface Response {status: number, message: string}

// let res:Response
export const checkCookie = async (x?: any) => {
  try {
    const res = await axios.get('/api/checkCookie')
    return res
  } catch (err) {
    throw new Error(err)
  }
  
  // if (x.value.id == "") {
  //   store.dispatch('loginWithCookie')
  // }
}