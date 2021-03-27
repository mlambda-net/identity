const { defaults } = require('jest-config');

module.exports = {
  verbose: true,
  roots: ['<rootDir>/test'],
  moduleFileExtensions: [...defaults.moduleFileExtensions, 'js', 'jsx'],
  testPathIgnorePatterns: ['./node_modules/'],
  moduleNameMapper: {
    '@mlambda-net/core/(.*)$': '<rootDir>/src/packages/$1',
  },

  collectCoverageFrom: [
    '**/*.{js,jsx}',
    '!**/node_modules/**',
    '!**/vendor/**',
  ],
  transform: {
    '^.+\\.(js|jsx)$': 'babel-jest',
  },

  snapshotResolver: './config/snapshot_resolver_test',
};
