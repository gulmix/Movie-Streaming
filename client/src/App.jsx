import './App.css'
import Home from './components/home/Home'
import Header from './components/header/Header'
import Register from './components/register/Register'
import Login from './components/login/Login'
import Recommended from './components/recommended/Recommended'
import { Route, Routes } from 'react-router-dom'
import Layout from './components/Layout'
import RequiredAuth from './components/RequiredAuth'

function App() {
  return (
    <>
      <Header/>
      <Routes path="/" element={<Layout/>}>
        <Route path='/' element={<Home/>} />
        <Route path='/register' element={<Register/>} />
        <Route path='/login' element={<Login/>} />
        <Route element={<RequiredAuth/>}>
          <Route path='/recommended' element={<Recommended/>}></Route>
        </Route>
      </Routes>
    </>
  )
}

export default App
