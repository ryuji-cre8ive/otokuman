import axios from "axios";
import { InjectionKey } from "vue";
import { createStore, useStore as baseUseStore, Store } from "vuex";

export interface User {
  id: string,
  name: string,
  password: string
}

// export interface State {
//   users: User[];
// }

// InjectionKeyを定義する

export const key: InjectionKey <Store<User>> = Symbol()

export const store = createStore<User>({
  state: {
    id: '',
    name: '',
    password: ''
  },
  getters: {
    getMyProfile: (state) => {
      return state
    }
  },
  
  mutations: {
    login(state, user) {
      state.id = user.Id
      state.name = user.name
      state.password = user.password
    },
    loginWithCookie(state) {
      axios.get('/api/getUserDataWithCookie').then(res => {
        const user = res.data
        console.log(user)
        state.id = user.Id
        state.name = user.name
        state.password = user.password
      })
    }
  },
  actions: {
    login({ commit, state }, user: User) {
      commit('login', user )
    },
    loginWithCookie({ commit, state }) {
      commit('loginWithCookie')
    }
  },
})

export function useStore() {
  return baseUseStore(key)
}
type GlobalStateType = ReturnType<typeof useStore>
export const GlobalStateKey: InjectionKey<GlobalStateType> = Symbol(
  'GlobalState'
)
