export default interface ICheckUp {
  id: number,
  name: string,
  created_timestamp: Date,
  updated_timestamp: Date,
  disease: {
    id: number,
    name: string,
    sequence?: string
  },
  similarity: number,
  sample?: string,
  is_match: boolean
}
