import { BASE, useMock } from '.'

export type Tag = {
  id: string
  name: string
}

export type GetTagsRequest = {}
export type GetTagsResponse = Tag[]

export const getTags = async (req?: GetTagsRequest): Promise<GetTagsResponse> => {
  if (useMock) {
    const { getTagsMock } = await import('./mock')
    return getTagsMock(req)
  }

  const res = await fetch(`${BASE}/tags`)
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  const json: GetTagsResponse = await res.json()
  return json
}

export type PostTagRequest = {
  name: string
}
export type PostTagResponse = Tag

export const postTag = async (req: PostTagRequest): Promise<PostTagResponse> => {
  if (useMock) {
    const { postTagMock } = await import('./mock')
    return postTagMock(req)
  }

  const res = await fetch(`${BASE}/tags`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(req)
  })
  if (!res.ok) {
    throw new Error(res.statusText)
  }
  const json: PostTagResponse = await res.json()
  return json
}
