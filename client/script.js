const URL = 'http://192.168.14.5:8000'

const createContainer = filename => {
  const container = document.createElement('div')
  container.classList.add('p-2', 'rounded', 'shadow-sm', 'ms-4', 'mb-2')
  container.style.backgroundColor = '#eee'
  container.style.width = '350px'
  container.innerHTML = `📄 ${filename}`
  return container
}

const createControls = filename => {
  const div = document.createElement('div')
  div.className = 'mt-2'
  const controls = [
    {
      label: 'download',
      endpoint: `${URL}/files/${filename}`,
      color: 'secondary',
    },
    { label: 'watch', endpoint: `${URL}/files/${filename}`, color: 'success' },
  ]
  const buttons = controls.map(btn => {
    const a = document.createElement('a')
    a.classList.add('me-1', 'btn', 'btn-sm', `btn-${btn.color}`)
    a.href = btn.endpoint
    a.innerHTML = btn.label
    if (btn.label === 'download') a.download = filename
    return a
  })
  buttons.forEach(btn => div.appendChild(btn))
  return div
}

const createEntry = fileName => {
  const c = createContainer(fileName)
  c.appendChild(createControls(fileName))
  return c
}

const filesNames = fetch(`${URL}/list`).then(res => JSON.parse(res))

const root = document.getElementById('root')

filesNames.map(fname => createEntry(fname)).forEach(e => root.appendChild(e))
