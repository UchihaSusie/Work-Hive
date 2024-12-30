
# WorkHive - Multi-Platform Task Management App
WorkHive is a task management application designed to streamline task organization and collaboration within teams. Built using  React Native and Golang, this platform supports user authentication and offers different access levels for owners, admins, and employees. It can  focuses on efficient task organization and scalable data management and provide seamless experiences on iOS and Android devices.

## Features

### 1. **Collaborative Development**
- **User Authentication**: Implemented a secure, role-based login system using JWT and Redis for session management, reducing login wait time through multithreading and parallel data retrieval.
- **Role-Specific Interfaces**: Conditional rendering in React allows for customized user experiences based on access levels.
- **State Management**: Integrated Redux for effective state management, enhancing responsiveness during data fetching with loaders and spinners.
- **Scalable Data Storage**: Utilized MongoDB for flexible, document-based storage to manage dynamic task data such as priorities, tags, and attachments.
- **Cross-Platform Development**: Delivered seamless user experiences on both iOS and Android devices using React Native for the frontend and Golang with the Gin framework for the backend.
- **Analytics Productivity**: Features for tracking user productivity and generating actionable insights, Add calendar integration for task scheduling.

      
## Technologies Used
- **Frontend**: React Native
- **Backend**: Golang with the Gin framework
- **Database**: MongoDB
- **Session Management**: Redis

## How to Run the Project
1. Clone the repository:
   ```bash
   git clone https://github.com/UchihaSusie/WorkHive.git
   ```
2. Navigate to the project directory:
   ```bash
   cd WorkHive
   ```
3. Install the necessary packages:
   ```bash
    npm install
4. Set up environment variables (create a .env file in the root directory):
   ```bash makefile
    MONGODB_URI=mongodb://localhost:27017
    REDIS_ADDR=localhost:6379
5. Start the application:
   ```bash
    npm start
6. Open your emulator or connect a physical device to test the application.

## Usage
Navigate to the login page and enter your credentials.
Access the platform based on your role (owner, admin, employee) and start managing tasks.

## Contribution Guidelines
Contributions are welcome! Feel free to fork the repository and submit a pull request.







