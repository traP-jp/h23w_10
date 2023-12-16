import { BASE, useMock } from '.'

export type UserType = 'trap' | 'external'

export type User = {
  id: string
  name: string
  icon_url: string
  userType: UserType
}


export type GetMeResponse = User

export const getMe = async (): Promise<GetMeResponse> => {
  if (useMock) {
    const { getMeMock } = await import('./mock')
    return getMeMock()
  }

  const res = await fetch(`${BASE}/users/me`)
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  const json: GetMeResponse = await res.json()
  return json
}

export type GetUserRequest = {
  id: User['id']
}
export type GetUserResponse = User

export const getUser = async (req: GetUserRequest): Promise<GetUserResponse> => {
  if (useMock) {
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
