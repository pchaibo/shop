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
//get address
export function getaddress() {
  return request({
    url: 'api/tronaddess',
    method: 'get',
  })
}

export function useradd(data) {
  return request({
    url: 'v1/user/add',
    method: 'post',
    data
  })
}

//site
export function SiteList(query) {
  return request({
    url: 'v1/site/list',
    method: 'get',
    params: query
  })
}
export function siteadd(data) {
  return request({
    url: 'v1/site/add',
    method: 'post',
    data
  })
}
export function Sitedel(id) {
  return request({
    url: 'v1/site/del',
    method: 'get',
    params: { id }
  })
}

