import request from '@/utils/request'
//get address
export function paytronlist() {
  return request({
    url: 'v1/paytron/list',
    method: 'get',
  })
}

