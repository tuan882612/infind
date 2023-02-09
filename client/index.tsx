import ReactDOM from 'react-dom/client'
import './src/assets/styles/index.css'
import App from './src/app'
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import Routing from './src/routes'

ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
).render(
  <App/>
)

window.sessionStorage.setItem('login','false')