import { useState } from "react";
import perfil from "../assets/usuario.png"
const IssueIdentity = () => {

        const [formData, setFormData] = useState({
          nombre: '',
          apellido: '',
          fechaNacimiento: '',
          numeroCarne: '',
          direccion: '',
          telefono: '',
          email: '',
        });
      
        const handleChange = (e) => {
          setFormData({ ...formData, [e.target.name]: e.target.value });
        };
      
        const handleSubmit = async (e) => {
          e.preventDefault();
          try {
            const response = await fetch('URL_DE_TU_API', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
              },
              body: JSON.stringify(formData),
            });
            if (response.ok) {
              alert('Registro exitoso');
              // Puedes añadir lógica adicional aquí después de enviar los datos
            } else {
              alert('Error al registrar');
            }
          } catch (error) {
            console.error('Error al registrar:', error);
            alert('Error al registrar');
          }
        };

  return (
    <div className="mt-20">
        <h2 className="text-3xl sm:text-4xl lg:text-5xl mt-10 lg:mt-10 tracking wide text-center">
            Test the Indentity Issue  
            <span className="flex justify-center bg-gradient-to-r from-sky-800 to-sky-500 text-transparent bg-clip-text">
                Smart Contract
            </span> 
        </h2>
        <div className="flex flex-wrap justify-center mt-10">
            <div className="p-2 w-full lg:w-1/2">
            formulario
            </div>
            <div className="p-2 w-full lg:w-1/2 items-center">
                <img src={perfil} type="image/png" alt="profile" className="w-32"/>
                <h5>ID:</h5>
                <h5>Name:</h5>
                <h5>Last Name:</h5>
                <h5>Faculty:</h5>
                <h5>Address:</h5>
                <h5>University:</h5>
                <h5>Semester:</h5>
            </div>
        </div>
    </div>
  )
}

export default IssueIdentity