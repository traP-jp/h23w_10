export type Question = {
  id: string
  userId: User['id']
  title: string
  content: string
  createdAt: Date
  tags: Tag[]
  answers: Answer[]
  status: QuestionStatus
}

export type Answer = {
  id: string
  userId: User['id']
  questionId: Question['id']
  content: string
  createdAt: Date
}

export type Tag = {
  id: string
  name: string
}

export type User = {
  id: string
  name: string
  iconURL: string
  userType: UserType
}

export type QuestionStatus = 'open' | 'closed'

export type UserType = 'trap' | 'external'
