import { Employee } from '@/types/employee'
import React from 'react'

export default function getEmployees(): Promise<Employee[]> {
  return Promise.resolve([
    {
      id: 1,
      registration_number: '123456',
      name: 'John Doe',
      work_unit: 'Work Unit 1',
      position: 'Position 1',
      status: 1,
      phone: '1234567890',
      email: 'cJtTt@example.com'
    }
  ])
}
