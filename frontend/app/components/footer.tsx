export default function Footer() {
    return(
        <footer className="bg-white text-black py-6 shadow-inner">
        <div className="container mx-auto flex justify-between items-center">
          <div>&copy; 2025 My Website. All rights reserved.</div>
          <div>
            <ul className="flex space-x-4">
              <li><a href="#" className="hover:text-indigo-600">Privacy Policy</a></li>
              <li><a href="#" className="hover:text-indigo-600">Terms of Service</a></li>
            </ul>
          </div>
        </div>
      </footer>
    )
}