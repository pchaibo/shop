import request from '@/utils/request'

// 取用户list
export function fetchList(query) {
  return request({
    url: 'v1/admin/list',
    method: 'get',
    params: query
  })
}
// 删除
export function userdel(id) {
  return request({
    url: 'v1/admin/del',
    method: 'get',
    params: { id }
  })
}
export function useradd(data) {
  return request({
    url: 'v1/admin/add',
    method: 'post',
    data
  })
}
