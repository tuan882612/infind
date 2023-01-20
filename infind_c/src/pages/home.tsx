import '../assets/styles/home.css'
import image1 from '../assets/images/man on computer.png'
import { Button, Typography } from '@mui/material'
import { Box } from '@mui/system'
import { useNavigate } from 'react-router-dom'

const Home = () => {
  const navigate = useNavigate();

  return (
    <>
      <Box className='midBody'>
        <Typography marginBottom='2vh' fontSize='20px'>
          Faster and better way to find opportunities.
        </Typography>
        <Button
          sx={{
            width:'150px',
            height:'40px',
            color:'white',
            bgcolor:'#75a9f9',
            borderRadius:'12px',
            '&:hover': { backgroundColor: '#a8c6fa' }
          }}
          onClick={() => navigate('/login')}
        >
          Get Started
        </Button>
      </Box>
      <img className='image1' src={image1}/>
    </>
  )
}

export default Home;