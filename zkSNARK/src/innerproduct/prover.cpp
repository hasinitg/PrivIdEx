//
// Created by Hasini Gunasinghe on 2/8/18.
//

#include <libsnark/gadgetlib2/adapters.hpp>
#include <libsnark/gadgetlib2/gadget.hpp>
#include <libsnark/gadgetlib2/integration.hpp>
#include "generator.hpp"
#include <sstream>
#include <type_traits>

#include <libff/common/profiling.hpp>
#include "r1cs_examples.hpp"
#include <libsnark/zk_proof_systems/ppzksnark/r1cs_ppzksnark/r1cs_ppzksnark.hpp>

namespace libsnark {

    r1cs_example<libff::Fr<libff::default_ec_pp>> generateInnerProductR1CSWithInputs(const size_t size) {

        typedef libff::Fr<libff::default_ec_pp> FieldT;

        gadgetlib2::initPublicParamsFromDefaultPp();

        gadgetlib2::GadgetLibAdapter::resetVariableIndex();

        auto pb = gadgetlib2::Protoboard::create(gadgetlib2::R1P);
        gadgetlib2::VariableArray A(size, "A");
        gadgetlib2::VariableArray B(size, "B");
        gadgetlib2::Variable result("result");
        auto g = gadgetlib2::InnerProduct_Gadget::create(pb, A, B, result);
        g->generateConstraints();

        for (size_t k = 0; k < size; ++k)
        {
            pb->val(A[k]) = std::rand() % 2;
            pb->val(B[k]) = std::rand() % 2;
        }
        g->generateWitness();

        r1cs_constraint_system<FieldT> cs = get_constraint_system_from_gadgetlib2(*pb);

        const r1cs_variable_assignment<FieldT> fullAssignment = get_variable_assignment_from_gadgetlib2(*pb);
        const r1cs_primary_input<FieldT> primary_input(fullAssignment.begin(), fullAssignment.begin() + cs.num_inputs());
        const r1cs_auxiliary_input<FieldT> auxiliary_input(fullAssignment.begin() + cs.num_inputs(), fullAssignment.end());

        return r1cs_example<FieldT>(cs, primary_input, auxiliary_input);
    }

    template <typename ppT>
    r1cs_ppzksnark_proof<ppT> generateProofForR1CS(const r1cs_ppzksnark_proving_key<ppT> &pk,
                                                  const r1cs_ppzksnark_primary_input<ppT> &primary_input,
                                                  const r1cs_ppzksnark_auxiliary_input<ppT> &auxiliary_input){
        return r1cs_ppzksnark_prover<ppT>(pk, primary_input, auxiliary_input);

    }
}