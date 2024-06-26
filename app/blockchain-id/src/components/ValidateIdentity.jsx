import { useState } from "react";
import perfil from "../assets/perfil.png"
import QRCode from 'qrcode.react';
import {
  Card,
  Input,
  Checkbox,
  Button,
  Typography,
  CardBody,
  CardFooter,
} from "@material-tailwind/react";

const ValidateIdentity = () => {

  const [inputValue, setInputValue] = useState('');

  const handleInputchange = (event) =>{
    setInputValue(event.target.value);
  };


  return (
    <div className="mt-20 border-t border-neutral-800">
        <h2 className="text-3xl sm:text-4xl lg:text-5xl mt-10 lg:mt-10 tracking wide text-center">
            Test the Indentity Validate  
            <span className="flex justify-center bg-gradient-to-r from-sky-800 to-sky-500 text-transparent bg-clip-text">
                Smart Contract
            </span> 
        </h2>
        <div className="flex flex-wrap items-center justify-center mt-10">
            <div className="p-2 w-full lg:w-1/2">
      <Card color="transparent" shadow={false} className="w-full flex items-center">
      <Typography color="gray" className="mt-1 font-normal">
        Nice to meet you! Enter your id to validate your identity.
      </Typography>
      <form  className="mt-8 mb-2 w-80 max-w-screen-lg sm:w-96">
        <div className="mb-1 flex flex-col gap-6">
          <Typography variant="h6" color="blue-gray" className="-mb-3 text-sky-500">
            Your ID
          </Typography>
          <Input
            type= "text"
            value={inputValue}
            onChange={handleInputchange}
            size="lg"
            placeholder="ex: 201820547"
            className=" !border-t-blue-gray-200 focus:!border-t-gray-900"
            labelProps={{
              className: "before:content-none after:content-none",
            }}
          />
          <div className="flex justify-center">
          <QRCode value={inputValue} size={256} className=" p-2 shadow-sm shadow-sky-500/50 rounded"/>
          </div>
          <Typography color="blue-gray" className="-mb-3 text-sky-500 text-center text-sm">
            Scan QR to Validate
          </Typography>
          <Typography color="blue-gray" className="-mb-3 text-center text-sm">
            or
          </Typography>
        </div>
        <Button type="submit" className="mt-6 bg-gradient-to-r from-sky-500 to-sky-800" fullWidth>
          Verify Identity
        </Button>
      </form>
    </Card>
    </div>
    <div className="p-2 w-full lg:w-1/2 mt-0 flex justify-center text-center">
    <Card className="shadow-lg shadow-sky-500/100 rounded-lg w-96 p-5">
      <CardBody>
        <div className="w-full flex justify-center">
        <img src={perfil} alt="default-profile-picture" className="w-44"/>
        </div>
        <Typography className="mb-2 mt-3 text-2xl">
          Data on Blockchain
        </Typography>
        <Typography>
            ID:
        </Typography>
        <Typography>
            ID:
        </Typography> 
        <Typography>
            ID:
        </Typography>
        <Typography>
            ID:
        </Typography>
        <Typography>
            ID:
        </Typography>
        <Typography>
            ID:
        </Typography>
        <Typography>
            ID:
        </Typography>
      </CardBody>
      <CardFooter className="pt-0">
      </CardFooter>
    </Card>
    </div>
        </div>
    </div>
  )
}

export default ValidateIdentity
