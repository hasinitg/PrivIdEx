//
// Created by Hasini Gunasinghe on 2/8/18.
//

#ifndef ZKSNARK_VERIFIER_HPP
#define ZKSNARK_VERIFIER_HPP

#include <libsnark/zk_proof_systems/ppzksnark/r1cs_ppzksnark/r1cs_ppzksnark.hpp>
#include "r1cs_examples.hpp"

namespace libsnark{

    template <typename ppT>
    const bool verifyProofOnR1CS(const r1cs_ppzksnark_verification_key<ppT> &vk,
                                 const r1cs_ppzksnark_primary_input<ppT> &primary_input,
                                 const r1cs_ppzksnark_proof<ppT> &proof);
}

#include "verifier.cpp"
#endif //ZKSNARK_VERIFIER_HPP
