import axios from './axios'

export interface StartGameResponse{
  uuid: string
}

export interface AskGameResponse{
  question: string
  answer: string
  hint: string
  gameend: boolean
}

export const postCreateGame = (soupID: number): Promise<{ data: StartGameResponse }> => {
  return axios.post('/game/create',
    {
      soup_id: soupID
    }
  )
}

export const postStartGame = (sessionID: string): Promise<{ data: {} }> => {
  return axios.post(`/game/${sessionID}/start`)
}

export const postAskGame = (sessionID: string, question: string, needHint: boolean): Promise<{ data: AskGameResponse }> => {
  return axios.post(`/game/${sessionID}/ask`,
    {
      question: question,
      need_hint: needHint
    }
  )
}

export const deleteEndGame = (sessionID: string) => {
  return axios.delete(`/game/${sessionID}/end`)
}
