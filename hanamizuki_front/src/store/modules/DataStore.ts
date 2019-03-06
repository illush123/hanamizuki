import axios from 'axios';

const state = {
  userProgram: {},
  output: {},
  test: {},
};

const getters = {};

const mutations = {
  /* payload = {code: 'code....'} */
  setUserProgram(state: any, payload: any): any {
    state.userProgram = payload.code;
  },

  /* payload = {output: 'output...'} */
  setOutput(state: any, payload: any): any {
    state.code = payload.output;
  }
};

const actions = {
  // stateのuserprogramに実行するコードを保存
  uploadUserProgram(context: any, payload: object): void {
    context.commit('setUserProgram', payload);
    console.log(state.userProgram);
    
    /* サーバにuserが書いたコードをポスト */
    axios.post('/api/v1/vmiss/upload', {
      'code': state.userProgram
    })
    .then ((response) => {
      console.log(response);
    })
    .catch((error) => {
      console.log(error);
    });
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
