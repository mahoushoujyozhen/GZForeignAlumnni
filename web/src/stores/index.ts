import { defineStore } from 'pinia'
import piniaPluginPersist from 'pinia-plugin-persist'

export const useUserStore = defineStore({
  id: 'app-user',
  state: () => ({
    userID:0,
    username:"",
    user_auth:false,
    admin_auth:false,
  }),
  persist: {
    enabled: true,
  },
  getters: {
    // doubleCount: (state) => state.counter * 2
    // getUserID() :number {
    //   return this.userID
    // },
  
  },
  actions: {
    increment() {
      this.counter++
    },
    getUserName() : string{
      return this.username
    },
    getUserID() :number {
      return this.userID
    },
    setUserID(userID){
      this.userID = userID;
    },
    setUsername(username){
      this.username = username
    }

  }
})
