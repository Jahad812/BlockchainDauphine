# EX 1 - Part A
Account has 131506461761462800 Wei, but it may be incorrect due to floating-point precision issues in JavaScript. Since JavaScript uses 64-bit floating-point numbers, large integers like Wei can lose precision when converted, leading to potential discrepancies.

---

# EX 1 Part B 
## SPDX Initiative 

The SPDX (Software Package Data Exchange) is a project by the Linux Foundation that creates a standard way to list open-source licenses. It helps developers understand and share software legally.
License Choices:
1.	DeFi Project: Use GNU General Public License (GPL-3.0). This ensures that all modified and shared versions remain open-source, which is good for community-driven finance projects.
2.	Google Cloud Web3: Use Mozilla Public License (MPL-2.0). This allows Google to open-source some libraries while keeping the core software proprietary.

---

## EX 1 Part C

A Virtual Machine (VM) is a software-based computer that runs programs like a physical machine. It provides an isolated environment where smart contracts execute securely, ensuring consistency across all nodes in a blockchain network. The Ethereum Virtual Machine (EVM) and Solana Virtual Machine (SVM) are examples of blockchain-specific VMs that process transactions and enforce smart contract rules.
The EVM is widely adopted, supporting multiple chains like Ethereum, BNB Chain, and Polygon. It benefits from a large developer community, extensive tooling (Hardhat, Foundry, Remix), and strong network effects. However, EVM transactions are often slower and more expensive due to Ethereum's congestion and reliance on the Ethereum gas model.
The SVM, used by Solana, is optimized for speed and parallel execution. Unlike the EVM, which processes transactions sequentially, the SVM can handle thousands of transactions in parallel, leading to lower costs and higher throughput. However, Solana's ecosystem is smaller, and its developer community is less mature than Ethereum’s, which impacts tooling availability and long-term stability.
Community Metrics from the Developer Report:

•	Monthly active developers: Ethereum leads with the most active developers, but Solana has one of the fastest-growing communities.

•	Total Commits: EVM-based projects have a steady flow of commits, while Solana saw sharp increases but also fluctuations.


Conclusion
The EVM has a more stable and well-supported developer community, making it a safer choice for long-term projects. The SVM offers superior performance but is still in the early stages of adoption. Future improvements in Solana’s ecosystem could make it more competitive.

