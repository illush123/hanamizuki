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
  },

  /* codeforceのApi叩いてテスト payload = {data: {...}}*/
  setTest(state: any, payload: any): any {
    state.test = payload.data;
  }
};

const actions = {
  // test
  async test(context: any) {
    //context.commit('setUserProgram', {'code': program});
    await axios
      .get('https://codeforces.com/api/user.info?handles=ryo1126')
      .then(response => {
        if (response.data.status === 'OK') {
          context.commit('setTest', {'data': response.data});
        }
      });
  },

  // stateのuserprogramに実行するコードを保存
  uploadUserProgram(context: any, payload: object): void {
    context.commit('setUserProgram', payload);
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
