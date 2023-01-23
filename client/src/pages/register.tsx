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
  Stack,
  Box
} from "@mui/material";
import { VisibilityOff, Visibility } from "@mui/icons-material";
import { useForm } from "react-hook-form";


const Register = () => {
  const [values, setValues] = useState<user>({
    email: '', 
    password: '',
    cPassword: ''
  })

  const [psw, setPsw] = useState<boolean>(false)

  const [cpsw, setCpsw] = useState<boolean>(false)

  const handleChange = (props: string) => (event: ChangeEvent<HTMLInputElement>): void => {
    const value = event.target.value
    setValues({...values, [props]: value})
  }

  const changeVis = (): void => setPsw(!psw)

  const changeCvis = (): void => setCpsw(!cpsw)

  const { handleSubmit } = useForm()

  const handleForm = (): void => {
    console.log(values)
  }

	return (
    <Box
      display='flex'
      justifyContent='center'
      marginTop='12vh'
    >
      <Box 
        sx={{
          bgcolor:'white',
          width:'35vh',
          height:'50vh',
          borderRadius:'18px'
        }}
      >
        <form onSubmit={handleSubmit(handleForm)}>
          <Typography
            color='#75A9F9'
            fontSize='30px'
            marginTop='3.5vh'
            align='center'
          >
            Create an Account
          </Typography>

          <Stack alignItems='center'>
            <FormControl sx={{mt:'3vh', width:'250px'}}>
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

            <FormControl sx={{mt:'2.5vh', width:'250px'}}>
              <InputLabel>confirm password</InputLabel>
              <FilledInput
                type={cpsw ? "text":"password"}
                value={values.cPassword}
                onChange={handleChange('cPassword')} 
                endAdornment={
                  <InputAdornment position="end">
                    <IconButton onClick={changeCvis}>
                      {cpsw ? <VisibilityOff/>:<Visibility/>}
                    </IconButton>
                  </InputAdornment>
                }
              />
            </FormControl>
            
            <Typography fontSize='13px' color='grey'>
              *password must be 8 characters long 
            </Typography>
            <FormButton name='Create' top='2vh'/>
          </Stack>
        </form>
      </Box>
    </Box>
	);
}

export default Register;