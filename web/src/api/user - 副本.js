import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/v1/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/v1/info',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/v1/logout',
    method: 'post'
  })
}
// 取用户list
export function fetchList(query) {
  return request({
    url: 'v1/user/list',
    method: 'get',
    params: query
  })
}
// 删除
export function userdel(id) {
  return request({
    url: 'v1/user/del',
    method: 'get',
    params: { id }
  })
}
export function useradd(data) {
  return request({
    url: 'v1/user/add',
    method: 'post',
    data
  })
}
