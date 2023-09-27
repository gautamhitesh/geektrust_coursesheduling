Your job is to build a simple command line application, which does the following:
 
Add course offering
 A course offering has course title, instructor and date. 
 It should also contain a minimum & maximum number of employees for the course offering. 
Register for the course
  Employees can register for the courses. 
 If no. of employees registered for the course has reached the maximum, the result will be COURSE_FULL_ERROR. 
 Otherwise, result of registration will be ACCEPTED. 
Cancel registration
 Employees can cancel their registration until the course allotment is completed. 
Course allotment
 This feature allots employees to course offering, before the course offering date. 
 It should print a list of all the employees with their details along with their final course allotment status (Registration Number, Employee Name, Email, Course Offering ID, Course Name, Instructor, Date, Final Status). The list should be sorted based on the Registration number. 
 If sufficient registrations are not received then the course offering itself gets cancelled. 
 The employees who have registered will get confirmed unless the minimum number of registrations is not received. 
 Even if the course offering gets canceled due to the minimum number of employees not registered, the list should be printed. 