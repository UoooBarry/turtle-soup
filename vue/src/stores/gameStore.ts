import { defineStore } from 'pinia'
import { ref, type Ref } from 'vue'
import { postCreateGame, postStartGame, postAskGame, type AskGameResponse, deleteEndGame } from '@/api/game.api'

export enum GameState {
  Picking = 'PICKING',
  Gaming = 'GAMING',
  Finished = 'FINISHED',
}

export const useGameStore = defineStore('game', () => {
  const state: Ref<GameState> = ref(GameState.Picking)
  const sessionID = ref('')
  const gameAnswer = ref('')

  const setState = (newState: GameState) => {
    state.value = newState
  }

  const createGame = (soupID: number) => {
    postCreateGame(soupID).then(({ data }) => {
      state.value = GameState.Gaming
      sessionID.value = data.uuid
    })

    return Promise.reject(false)
  }

  const startGame = (): Promise<boolean> => {
    return postStartGame(sessionID.value).then(() => {
      return true
    })
      .catch(() => {
        state.value = GameState.Picking
        return false
      })
      .finally(() => {
        return true
      })
  }

  const askGame = async (question: string): Promise<AskGameResponse> => {
    const { data } = await postAskGame(sessionID.value, question)
    if (data.gameend === true) {
      state.value = GameState.Finished
      gameAnswer.value = data.answer
    }
    return data
  }

  const endGame = async (): Promise<boolean> => {
    try {
      await deleteEndGame(sessionID.value)
      sessionID.value = ''
      return Promise.resolve(true)
    } catch(e) {
      console.error(e)
      return Promise.reject(false)
    } finally {
      state.value = GameState.Picking
    }
  }

  return { state, setState, createGame, startGame, askGame, endGame, gameAnswer }
})

