

export default function Header() {
    return(
        <header className="bg-white text-black shadow-md py-4">
        <div className="container mx-auto flex justify-between items-center">
          <div className="text-2xl font-black text-indigo-800">HuCap</div>
          <nav>
            <ul className="flex space-x-6">
              <li><a href="#" className="hover:text-indigo-600">Home</a></li>
              <li><a href="#" className="hover:text-indigo-600">About</a></li>
              <li><a href="#" className="hover:text-indigo-600">Services</a></li>
              <li><a href="#" className="hover:text-indigo-600">Contact</a></li>
            </ul>
          </nav>
        </div>
      </header>
    )
}