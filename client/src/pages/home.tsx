import '../assets/styles/home.css'
import image1 from '../assets/images/man on computer.png'
import logo from '../assets/images/infind-logo.png'
import { Button, Stack, Typography } from '@mui/material'
import { Box } from '@mui/system'
import { useNavigate } from 'react-router-dom'

const style = {
  width:'150px',
  height:'40px',
  color:'white',
  bgcolor:'#75a9f9',
  borderRadius:'12px',
  '&:hover': { backgroundColor: '#a8c6fa' }
}

const Home = () => {
  const navigate = useNavigate();

  return (
    <>
      <img className='logo' src={logo}/>
      <Stack direction='row' alignItems='center'>
        <Box
          ml='25%'
          textAlign='center'
        >
          <Typography marginBottom='2vh' fontSize='20px'>
            Faster and better way to find opportunities.
          </Typography>
          <Button
            sx={style}
            onClick={() => navigate('/login')}
          >
            Get Started
          </Button>
        </Box>
        <img className='image1' src={image1}/>
      </Stack>
    </>
  )
}

export default Home;