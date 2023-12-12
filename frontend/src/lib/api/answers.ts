import { BASE } from '.'
import type { Question } from './questions'
import type { User } from './users'

export type Answer = {
  id: string
  userId: User['id']
  questionId: Question['id']
  content: string
  createdAt: Date
}

export type PostAnswerRequest = Omit<Answer, 'id' | 'createdAt'>
export type PostAnswerResponse = Answer

export const postAnswer = async (req: PostAnswerRequest): Promise<PostAnswerResponse> => {
  if (import.meta.env.DEV) {
    const { postAnswerMock } = await import('./mock')
    return postAnswerMock(req)
  }

  const res = await fetch(`${BASE}/questions/${req.questionId}/answers`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(req)
  })
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  const json: PostAnswerResponse = await res.json()
  return json
}

// TODO: PUT /questions/:questionId/answers/:answerId
