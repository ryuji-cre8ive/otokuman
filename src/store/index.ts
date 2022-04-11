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
    }
  },
  actions: {
    login({ commit, state }, user: User) {
      commit('login', user )
    }
  },
})

export function useStore() {
  return baseUseStore(key)
}
