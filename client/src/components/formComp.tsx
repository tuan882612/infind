import { Button, FormControl } from "@mui/material"
import { ChangeEvent, FunctionComponent } from "react"
import cButton from "../types/cButton"

const FormButton: FunctionComponent<cButton> = ({name, top, bottom}) => {

  return (
    <Button
      sx={{
        bgcolor:'#75A9F9',
        color:'white',
        mt:top,
        mb:bottom,
        width:'15vh',
        borderRadius:'10px',
        '&:hover':{backgroundColor:'#c3d8f7'}
      }}
      type="submit"
    >
      {name}
    </Button>
  )
}

export { FormButton }