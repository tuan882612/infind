import { ChangeEvent, FormEventHandler, useState } from "react";

import user from "../types/user";
import { FormButton } from "../components/formComp";

import { 
  FormControl, 
  InputAdornment, 
  FilledInput,
  InputLabel,
  IconButton,
  Typography,
  Link,
  Stack,
  Box,
  Button
} from "@mui/material";
import { VisibilityOff, Visibility } from "@mui/icons-material";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const styleBase = { display:'flex', justifyContent:'center', mt:'25vh' }

const styleHeader = { color:'#75A9F9', fontSize:'38px', mt:'3.5vh'}

const styleMain = {
  bgcolor:'white',
  width:'35vh',
  height:'50vh',
  textAlign: 'center',
  borderRadius:'18px',
}

const styleForm = [
  {mt:'4vh', width:'250px'},
  {mt:'2.5vh', width:'250px'}
]

const styleForgotPsw = { color:'grey', fontSize:'13px', ml:'55%' }

const Login = () => {  
  const navigate = useNavigate()
  const { handleSubmit } = useForm()

  const [values, setValues] = useState<user>({
    username: '',
    password: ''
  })

  const [psw, setPsw] = useState<boolean>(false)

  const handleChange = (props: string) => 
    (event: ChangeEvent<HTMLInputElement>): void => {
    const value = event.target.value
    setValues({...values, [props]: value})
  }

  const changeVis = (): void => setPsw(!psw)

  const handleForm = async (): Promise<void> => {
    console.log(values)
    await axios.get(
      'http://localhost:1000/api/v1/login'+
      `?username=${values.username}&password=${values.password}`
    ).then(res => {
      if (res.data) {
        window.sessionStorage.setItem('login','true')
        window.sessionStorage.setItem('username',values.username)
        navigate('/dashboard', { replace: true })
      } else {
        console.log('invalid')
      }
    }).catch(err => console.log(err))
  }

  return (
    <Box sx={styleBase}>
      <Box sx={styleMain}>
        <Typography sx={styleHeader}>
          Login
        </Typography>
        
        <form onSubmit={handleSubmit(handleForm)}>
          <Stack alignItems='center'>
            <FormControl sx={styleForm[0]}>
              <InputLabel>username</InputLabel>
              <FilledInput 
                type="text"
                value={values.username}
                onChange={handleChange('username')}
              />
            </FormControl>

            <FormControl sx={styleForm[1]}>
              <InputLabel>password</InputLabel>
              <FilledInput
                type={psw ? "text":"password"}
                value={values.password}
                onChange={handleChange('password')} 
                endAdornment={
                  <InputAdornment position="end">
                    <IconButton onClick={changeVis}>
                      {psw ? <VisibilityOff/>:<Visibility/>}
                    </IconButton>
                  </InputAdornment>
                }
              />
            </FormControl>
          </Stack>
          <Link sx={styleForgotPsw}
            href=''
          >
            forgot password?
          </Link>

          <Stack alignItems='center' spacing={1}>
            <FormButton name='Sign In' top='2vh'/>
            <Typography fontSize='17px' color='grey'>
              or 
            </Typography>
            <Link href='/register'>sign up</Link>
          </Stack>

        </form>
      </Box>
    </Box>
  );
}

export default Login;