import videob1  from "../assets/videob1.gif"
import videob2  from "../assets/videob2.gif"
const Home = () => {
  return (
    <div className="flex flex-col items-center mt-6 lg:mt-20">
        <h1 className="text-4xl sm:text-6xl lg:text-7xl text-center tracking-wide">
            The Descentralized Identity 
            <span className="bg-gradient-to-r from-sky-500 to-sky-800 text-transparent bg-clip-text">
                {" "}
                Prototype app is Here</span>
        </h1>
        <p className="mt-10 text-lg text-center text-neutral-500 max-w-4xl">
        In today's digital era, safeguarding our online identity is crucial, with blockchain technology offering a revolutionary solution to traditional vulnerabilities.
        </p>
        <div className="flex justify-center my-10">
            <a href="#" className="bg-gradient-to-r from-sky-500 to-sky-800 py-3 px-4 mx-3 rounded-md transform transition duration-500 hover:scale-125">
            Documentation
            </a>
            <a href="#" className="py-3 px-4 mx-3 rounded-md border transform hover:bg-emerald-300 transition duration-500 hover:scale-125">
                Dashboard
            </a>
        </div>
        <div className="flex mt-10 justify-center max-w-screen-lg" >
        <img className="rounded-lg w-1/2 border border-sky-800 shadow-sky-500 mx-2 my-4" src={videob1} type="image/gif" alt="blockchain1" />
        <img className="rounded-lg w-1/2 border border-sky-800 shadow-sky-500 mx-2 my-4" src={videob2} type="image/gif" alt="blockchain2" />
        </div>
    </div>
  )
}

export default Home