import { BASE } from '.'

export type UserType = 'trap' | 'external'

export type User = {
  id: string
  name: string
  iconURL: string
  userType: UserType
}

export type GetUserRequest = {
  id: User['id']
}
export type GetUserResponse = User

export const getUser = async (req: GetUserRequest): Promise<GetUserResponse> => {
  if (import.meta.env.DEV) {
    const { getUserMock } = await import('./mock')
    return getUserMock(req)
  }

  const res = await fetch(`${BASE}/users/${req.id}`)
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  const json: GetUserResponse = await res.json()
  return json
}

// TODO: PUT /users/:userId
