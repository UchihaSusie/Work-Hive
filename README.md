
# TaskFree - Multi-Platform Task Management App

TaskFree is a cross-platform task management application designed to provide seamless experiences on iOS and Android devices. Built using React Native and Golang, it focuses on efficient task organization and scalable data management.

## Features

### 1. **Collaborative Development**
- using **React Native** for mobile and **Golang** with the Gin framework for the backend.
- seamless user experience on both iOS and Android devices.
- Built a secure, role-based login system leveraging **Redis** for session management.
- Reduced login wait time by 30% through multithreading and parallel data retrieval.
- Integrated **MongoDB** for flexible, document-based storage of dynamic task data such as priorities, tags, and attachments.
- Enabled real-time updates and high scalability to handle diverse user needs.

## Technologies Used
- **Frontend**: React Native
- **Backend**: Golang with the Gin framework
- **Database**: MongoDB
- **Session Management**: Redis

## How to Run the Project
1. Clone the repository:
   ```bash
   git clone https://github.com/UchihaSusie/taskfree.git
   ```
2. Navigate to the project directory:
   ```bash
   cd taskfree
   ```
3. Install dependencies:
   ```bash
   npm install
   ```
4. Start the development server:
   ```bash
   npm start
   ```
5. Open your emulator or connect a physical device to test the application.

## Future Enhancements
- Add calendar integration for task scheduling.
- Implement push notifications for deadlines and reminders.
- Enhance analytics for user productivity tracking.

## Contribution Guidelines
Contributions are welcome! Feel free to fork the repository and submit a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments
Thanks to the team for their collaborative efforts and the open-source community for providing tools and frameworks that made this project possible.




# WorkHive

## Overview
WorkHive is a task management application designed to streamline task organization and collaboration within teams. Built using React, Node.js, Express.js, and MongoDB, this platform supports user authentication and offers different access levels for owners, admins, and employees.

## Features
- **User Authentication**: Secure login system implemented using JSON Web Tokens (JWT) for role-based access.
- **Role-Specific Interfaces**: Conditional rendering in React allows for customized user experiences based on access levels.
- **State Management**: Integrated Redux for effective state management, enhancing responsiveness during data fetching with loaders and spinners.
- **Performance Improvements**: Increased user interaction speed by approximately 30%, leading to improved overall satisfaction.
- **Deployment**: Successfully deployed in a robust cloud environment, ensuring scalability and reliability.

## Technologies Used
- **Frontend**: React
- **Backend**: Node.js, Express.js
- **Database**: MongoDB
- **State Management**: Redux
- **Authentication**: JSON Web Tokens (JWT)

## Installation
To run this project locally, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/YourUsername/WorkHive.git
   cd WorkHive
2. Install the necessary packages:
   ```bash
    npm install
4. Set up environment variables (create a .env file in the root directory):
   ```bash makefile
    MONGODB_URI=your_mongodb_uri
    JWT_SECRET=your_jwt_secret
5. Start the application:
   ```bash
    npm start
## Usage
Navigate to the login page and enter your credentials.
Access the platform based on your role (owner, admin, employee) and start managing tasks.

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.







