import { ChangeEvent, useState } from "react";

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
  Box
} from "@mui/material";
import { VisibilityOff, Visibility } from "@mui/icons-material";
import { useForm } from "react-hook-form";
// import { margin } from "@mui/system";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const [values, setValues] = useState<user>({
    email: '',
    password: ''
  })

  const [psw, setPsw] = useState<boolean>(false)

  const navigate = useNavigate()

  const handleChange = (props: string) => 
    (event: ChangeEvent<HTMLInputElement>): void => {
    const value = event.target.value
    setValues({...values, [props]: value})
  }

  const changeVis = (): void => setPsw(!psw)

  const { handleSubmit } = useForm()

  const handleForm = (): void => {
    if (values.email === 'root' && values.password === '12345') {
      window.sessionStorage.setItem('login','true')
      navigate('/dashboard')
    }
  }

  return (
    <Box
      sx={{
        display:'flex',
        justifyContent:'center',
        mt:'12vh'
      }}
    >
      <Box 
        sx={{
          bgcolor:'white',
          width:'35vh',
          height:'50vh',
          borderRadius:'18px',
        }}
      >
        <Typography 
          sx={{
            color:'#75A9F9',
            fontSize:'38px',
            mt:'3.5vh',
            textAlign:'center'
          }}
        >
          Login
        </Typography>
        
        <form onSubmit={handleSubmit(handleForm)}>
          <Stack alignItems='center'>
            <FormControl sx={{mt:'4vh', width:'250px'}}>
              <InputLabel>email</InputLabel>
              <FilledInput 
                type="text"
                value={values.email}
                onChange={handleChange('email')}
              />
            </FormControl>

            <FormControl sx={{mt:'2.5vh', width:'250px'}}>
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
          <Link 
            sx={{
              color:'grey', 
              fontSize:'13px',
              ml:'55%'
            }}
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