import simpleRestProvider from "ra-data-simple-rest";
const url = import.meta.env.VITE_SIMPLE_REST_URL

const unimpl = (resource, params) => {
  return new Promise((resolve, reject) => {
    console.log(resource)
    console.log(params)
    reject("unimplemented");
  });
}

const getOne = (resource, params) => {
  console.log(params)
  if (params.meta != undefined) {
    params.filter = {parent: params.meta.parent}
  }
  return new Promise((resolve, reject) => {
    call("GET", resource, params).then((resp) => {
      resp.id = 0
      resolve({data: resp})
    }).catch((e) => {
      reject(e)
    })
  })
}

const getList = (resource, params) => {
  let listFilter = []
  if (params.filter?.eq != undefined) {
    let eqFilter = []
    for (let key in params.filter.eq) {
      eqFilter.push(`${key} = ${params.filter.eq[key]}`)
    }
    listFilter.push(eqFilter.join(" AND "))
  }
  params.filter.listFilter = listFilter.join(" AND ")
  return new Promise((resolve, reject) => {
    call("GET", resource, params).then((data) => {
      const arr = data[resource]
      const dataArr = arr.map((o, i) => {
        o.id = o.displayName
      })
      resolve({data: arr, total: arr.length})
    }).catch((e) => {
      reject(e)
    })
  })
}

const call = (method, resource, params) => {
  console.log(params)
  console.log(resource)
  let parentPath = "";
  if (params.filter != undefined && params.filter.parent != undefined) {
    parentPath = params.filter.parent + "/"
  }
  let resourcePath = resource
  if (params.id != undefined) {
    resourcePath += `/${params.id}`
  }
  return new Promise((resolve, reject) => {
    fetch(`${url}/${parentPath}${resourcePath}`, {
      method: 'GET',
      headers: {
        'content-type': 'application/json;charset=UTF-8',
      },
    }).then((resp) => { return resp.json() }).then((data) => {
      resolve(data)
    }).catch((e) => {reject(e)})
  })
}

export const dataProvider = {
    create: unimpl,
    delete: unimpl,
    deleteMany: unimpl,
    getList: getList,
    getMany: unimpl,
    getManyReference: unimpl,
    getOne: getOne,
    update: unimpl,
    updateMany: unimpl,
};
