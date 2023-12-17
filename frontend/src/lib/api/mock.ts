// モックデータ/関数をまとめたファイルです! 本番環境でこれらの値・関数を使用しないように注意してください!

import type { Answer, PostAnswerRequest, PostAnswerResponse } from './answers'
import type {
  GetQuestionRequest,
  GetQuestionResponse,
  GetQuestionsRequest,
  GetQuestionsResponse,
  PostQuestionRequest,
  PostQuestionResponse,
  Question
} from './questions'
import type { GetTagsRequest, GetTagsResponse, PostTagRequest, PostTagResponse, Tag } from './tags'
import type { GetUserRequest, GetUserResponse, User } from './users'

const tags: Tag[] = [
  'Programming',
  'Sound',
  'CTF',
  'Python',
  'Graphic',
  'TypeScript',
  'ClipStudioPaint'
].map((name) => ({
  id: crypto.randomUUID(),
  name
}))

const randomChoice = <T>(arr: T[], count: number): T[] => {
  if (arr.length <= count) return arr
  const copy = [...arr]
  return copy.sort(() => Math.random() - 0.5).slice(0, count)
}
const sleep = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms))

const answers: Answer[] = []
const users: User[] = [
  {
    id: 'test-user-0',
    name: 'テストユーザー',
    icon_url: 'https://placehold.jp/150x150.png',
    userType: 'trap'
  }
]
const questions: Question[] = new Array(100)
  .fill({
    content: 'テストの質問です',
    createdAt: new Date(2023, 11, 13)
  } satisfies Partial<Question>)
  .map<Question>((question, i) => {
    const id = `test-question-${i}`
    return {
      ...question,
      id,
      title: `テストの質問${i}`,
      userId: crypto.randomUUID(),
      tags: randomChoice(tags, Math.floor(Math.random() * 3)),
      answers: new Array(Math.floor(Math.random() * 10)).fill({
        id: crypto.randomUUID(),
        questionId: id,
        content: 'テストの回答です',
        userId: crypto.randomUUID()
      } satisfies Omit<Answer, 'createdAt'>),
      status: Math.random() < 0.5 ? 'open' : 'closed'
    }
  })

/**
 * **モックAPI**です! 代わりに`getQuestions` (/src/lib/api/questions.ts)を使ってください。(開発環境では勝手にモックが使用されます)
 */
export const getQuestionsMock = async (req: GetQuestionsRequest): Promise<GetQuestionsResponse> => {
  await sleep(1000)
  const ret = questions.filter((question) => {
    let flag = true
    if (req.tag) {
      flag = (flag && question.tags?.some((tag) => tag.id === req.tag)) ?? false
    }
    if (req.status) {
      flag = flag && question.status === req.status
    }
    return flag
  })
  return {
    questions: ret.slice(req.offset, req.offset + req.limit),
    total: ret.length
  }
}

/**
 * **モックAPI**です! 代わりに`getQuestion` (/src/lib/api/questions.ts)を使ってください。(開発環境では勝手にモックが使用されます)
 */
export const getQuestionMock = async (req: GetQuestionRequest): Promise<GetQuestionResponse> => {
  await sleep(1000)
  const ret = questions.find((question) => question.id === req.id)
  if (!ret) {
    throw new Error('Not found')
  }
  return ret
}

/**
 * **モックAPI**です! 代わりに`postQuestion` (/src/lib/api/questions.ts)を使ってください。(開発環境では勝手にモックが使用されます)
 */
export const postQuestionMock = async (req: PostQuestionRequest): Promise<PostQuestionResponse> => {
  await sleep(1000)
  const question: Question = {
    ...req,
    id: crypto.randomUUID(),
    createdAt: new Date(),
    tags: req.tags
      .map((tag) => tags.find((t) => t.id === tag.id))
      .filter((tag): tag is Tag => tag !== undefined),
    answers: [],
    status: req.status ?? 'open'
  }
  questions.push(question)
  return question
}

// TODO: PUT /questions

/**
 * **モックAPI**です! 代わりに`postAnswer` (/src/lib/api/answers.ts)を使ってください。(開発環境では勝手にモックが使用されます)
 */
export const postAnswerMock = async (req: PostAnswerRequest): Promise<PostAnswerResponse> => {
  await sleep(1000)
  const answer: Answer = {
    ...req,
    id: crypto.randomUUID(),
    createdAt: new Date()
  }
  answers.push(answer)
  return answer
}

// TODO: PUT /questions/:questionId/answers/:answerId

/**
 * **モックAPI**です! 代わりに`getTags` (/src/lib/api/tags.ts)を使ってください。(開発環境では勝手にモックが使用されます)
 */
export const getTagsMock = async (_req?: GetTagsRequest): Promise<GetTagsResponse> => {
  await sleep(1000)
  return tags
}

/**
 * **モックAPI**です! 代わりに`postTag` (/src/lib/api/tags.ts)を使ってください。(開発環境では勝手にモックが使用されます)
 */
export const postTagMock = async (req: PostTagRequest): Promise<PostTagResponse> => {
  await sleep(1000)
  const tag: Tag = {
    id: crypto.randomUUID(),
    name: req.name
  }
  tags.push(tag)
  return tag
}

export const getUserMock = async (req: GetUserRequest): Promise<GetUserResponse> => {
  await sleep(1000)
  const user = users.find((user) => user.id === req.id)
  if (!user) {
    throw new Error('Not found')
  }
  return user
}

export const getMeMock = async (): Promise<GetUserResponse> => {
  await sleep(1000)
  const user = users[0]
  if (!user) {
    throw new Error('Not found')
  }
  return user
}
